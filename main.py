# -*- encoding : utf-8 -*-
import copy
import json
import os
import random
import re
import time

from khl import Bot, Message, Cert
from khl.card import Card, Types, Module, CardMessage, Element

from solution import get_solution

# webhook
# bot = Bot(cert=Cert(token='token', verify_token='verify_token'), port=3000,
#           route='/khl-wh')

# websocket
bot = Bot(token='token')

cache = {}


@bot.command(regex=r'(?:.|\/|。)(?:24开始)')
async def twenty_four_init(msg: Message):
    global cache
    if msg.author.bot:
        await msg.reply('暂不支持机器人游玩')
        return
    cache_id = f'{msg.ctx.guild.id}-{msg.ctx.channel.id}-{msg.author_id}'
    if cache_id not in cache:
        while True:
            cards = [random.randint(1, 13) for _ in range(4)]
            answer = get_solution(cards)
            if len(answer) != 0:
                break
        cache[cache_id] = {'cards': cards, 'time': 0.0, 'answer': answer}
        if len(cards) == 4:
            cache[cache_id]['original_cards'] = cards.copy()
        response = await msg.reply(
            f'来一把紧张刺激的 24 点！输入算式进行推导，输入「24退出」结束游戏\n(met){msg.author_id}(met) 现在你手上有：{cards}，怎么凑 24 点呢？'
        )
        cache[cache_id]['time'] = response['msg_timestamp']
    else:
        await msg.reply(f'24点游戏还没结束哦~')


@bot.command(regex=r'(?:.|\/|。)(?:24退出)')
async def twenty_four_exit(msg: Message):
    global cache
    if msg.author.bot:
        await msg.reply('暂不支持机器人游玩')
        return
    cache_id = f'{msg.ctx.guild.id}-{msg.ctx.channel.id}-{msg.author_id}'
    if cache_id not in cache:
        await msg.reply(f'没有正在进行的24点游戏')
    else:
        time_used = '%.2f' % ((msg.msg_timestamp - cache[cache_id]['time']) / 1000)
        answer = cache[cache_id]['answer']
        if len(answer) > 5:
            answer = '\n'.join(answer[:5]).replace('*', '\\*')
            answer += '\n答案仅显示前5种'
        else:
            answer = '\n'.join(answer[:5]).replace('*', '\\*')
        await msg.reply(f'24点游戏已退出, 这不再来一把？\n用时: {time_used}s\n{answer}')
        del cache[cache_id]


@bot.command(regex=r'[\d\+\-\\\*\/]+')
async def twenty_four_step(msg: Message):
    global cache
    content = msg.content.replace('\\*', '*')
    cache_id = f'{msg.ctx.guild.id}-{msg.ctx.channel.id}-{msg.author_id}'
    if cache_id not in cache:
        return
    if msg.author.bot:
        await msg.reply('暂不支持机器人游玩')
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
    time_used = '%.2f' % ((msg.msg_timestamp - cache[cache_id]['time']) / 1000)
    if len(cards) == 1 and cards[0] == 24:
        await msg.reply(f'你赢啦！\n用时: {time_used}s')
        del cache[cache_id]
        await add_list(msg.author_id, time_used)
    elif len(cards) == 1 and cards[0] != 24:
        answer = cache[cache_id]['answer']
        await msg.reply(f'你输啦！\n用时: {time_used}s\n{answer}')
        del cache[cache_id]
    else:
        await msg.reply(f'(met){msg.author_id}(met) 现在你手上有：{cards}，怎么凑 24 点呢？')


async def add_list(user_id, time_used):
    if not os.path.exists('top.json'):
        with open('top.json', 'w') as f:
            f.write(json.dumps({user_id: time_used}))
    else:
        with open('top.json', 'r') as f:
            data = json.loads(f.read())
        if len(data) == 0:
            data[user_id] = time_used
        elif len(data) < 10 and (user_id not in data):
            data[user_id] = time_used
        elif float(time_used) < float(data[list(data.keys())[len(data) - 1]]):
            if user_id not in data or (user_id in data and float(time_used) < float(data[user_id])):
                data[user_id] = time_used
        data = dict(sorted(data.items(), key=lambda x: float(x[1])))
        if len(data) > 10:
            count = 1
            data_new = {}
            for k, v in data.items():
                data_new[k] = v
                count += 1
                if count >= 11:
                    break
            data = data_new
        with open('top.json', 'w') as f:
            f.write(json.dumps(data))


async def get_list():
    if not os.path.exists('top.json'):
        return {}
    else:
        with open('top.json', 'r') as f:
            data = json.loads(f.read())
        return data


@bot.command(regex=r'(?:.|\/|。)(?:24排行榜)')
async def twenty_four_list(msg: Message):
    if msg.author.bot:
        await msg.reply('暂不支持机器人游玩')
        return
    d = await get_list()
    if len(d) != 0:
        text = ''
        count = 1
        for k, v in d.items():
            user = await msg.gate.request('GET', 'user/view', params={'user_id': k})
            name = f"{user['username']}#{user['identify_num']}"
            text += f'第{count}: {name} 用时: {v}s\n'
            count += 1
    else:
        text = '暂无数据'
    c = Card()
    c.theme = Types.Theme.WARNING
    c.append(Module.Header('24点前10排行榜'))
    c.append(Module.Section(Element.Text(content=text)))
    await msg.reply(CardMessage(c))


if __name__ == '__main__':
    bot.run()
