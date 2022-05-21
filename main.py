# -*- encoding : utf-8 -*-
import copy
import random
import re
import time
from khl import Bot, Message, Cert

# webhook
# bot = Bot(cert=Cert(token='token', verify_token='verify_token'), port=3000,
#           route='/khl-wh')

# websocket
bot = Bot(token='token')


# 解法来自于知乎用户 @曲晋云 在 https://zhuanlan.zhihu.com/p/37608401 评论区内的回答
class Solution:
    solutions = set()

    def point24(self, numbers):
        if len(numbers) == 1:
            if abs(eval(numbers[0]) - 24) < 0.00001:
                self.solutions.add(numbers[0])
        else:
            for i in range(len(numbers)):
                for j in range(i + 1, len(numbers)):
                    rest_numbers = [x for p, x in enumerate(numbers) if p != i and p != j]
                    for op in "+-*/":
                        if op in "+-*" or eval(str(numbers[j])) != 0:
                            self.point24(["(" + str(numbers[i]) + op + str(numbers[j]) + ")"] + rest_numbers)
                        if op == "-" or (op == "/" and eval(str(numbers[i])) != 0):
                            self.point24(["(" + str(numbers[j]) + op + str(numbers[i]) + ")"] + rest_numbers)

    def get_answer(self):
        return self.solutions

    def get_answer_top5_text(self):
        if len(self.solutions) == 0:
            return '无答案'
        answer = ''
        count = 1
        for i in self.solutions:
            answer += f'{i}\n'
            count += 1
            if count >= 6:
                break
        return answer.replace('*', '\\*')


cache = {}


@bot.command(regex=r'(?:24点)')
async def twenty_four_init(msg: Message):
    global cache
    cache_id = f'{msg.ctx.guild.id}-{msg.ctx.channel.id}-{msg.author_id}'
    if cache_id not in cache:
        cards = [random.randint(1, 13) for i in range(4)]
        cache[cache_id] = {'cards': cards, 'cards_source': cards.copy(), 'time': time.time()}
        await msg.reply(f'来一把紧张刺激的 24 点！输入算式进行推导，输入「24退出」结束游戏')
        await msg.reply(f'现在你手上有：{cards}，怎么凑 24 点呢？')
    else:
        await msg.reply(f'24点游戏还没结束哦~')


@bot.command(regex=r'(?:24退出)')
async def twenty_four_exit(msg: Message):
    global cache
    cache_id = f'{msg.ctx.guild.id}-{msg.ctx.channel.id}-{msg.author_id}'
    if cache_id not in cache:
        await msg.reply(f'没有正在进行的24点游戏')
    else:
        solution = Solution()
        solution.point24(cache[cache_id]['cards_source'])
        time_used = int(time.time() - cache[cache_id]['time'])
        await msg.reply(f'24点游戏已退出, 这不再来一把？\n用时: {time_used}s\n{solution.get_answer_top5_text()}')
        del cache[cache_id]
        del solution


@bot.command(regex=r'[\d\+\-\\\*\/]+')
async def twenty_four_step(msg: Message):
    global cache
    content = msg.content.replace('\\*', '*')
    cache_id = f'{msg.ctx.guild.id}-{msg.ctx.channel.id}-{msg.author_id}'
    if cache_id not in cache:
        return
    n_c = copy.deepcopy(cache[cache_id]['cards'])
    used = [int(i) for i in re.findall(r'\d+', content)]
    if 0 in map(lambda x: n_c.remove(x) if x in n_c else 0, used):
        await msg.reply('有错误！')
        return
    cards = n_c + [eval(content)]
    cards_new = []
    for i in cards:
        cards_new.append(int(i))
    cards = cards_new
    cache[cache_id]['cards'] = cards
    if len(cards) == 1 and cards[0] == 24:
        time_used = int(time.time() - cache[cache_id]['time'])
        await msg.reply(f'你赢啦！\n用时: {time_used}s')
        del cache[cache_id]
    elif len(cards) == 1 and cards[0] != 24:
        solution = Solution()
        solution.point24(cache[cache_id]['cards_source'])
        time_used = int(time.time() - cache[cache_id]['time'])
        await msg.reply(f'你输啦！\n用时: {time_used}s\n{solution.get_answer_top5_text()}')
        del cache[cache_id]
        del solution
    else:
        await msg.reply(f'现在你手上有：{cards}，怎么凑 24 点呢？')


if __name__ == '__main__':
    bot.run()
