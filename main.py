# -*- encoding : utf-8 -*-
import copy
import json
import random
import re
from khl import Bot, Message, Cert

# webhook
# bot = Bot(cert=Cert(token='token', verify_token='verify_token'), port=3000,
#           route='/khl-wh')

# websocket
bot = Bot(token='token')

cache = {}


@bot.command(regex=r'(?:24点)')
async def twenty_four_init(msg: Message):
    cache_id = f'{msg.ctx.guild.id}-{msg.ctx.channel.id}-{msg.author_id}'
    if cache_id not in cache:
        cards = [random.randint(1, 13) for i in range(4)]
        cache[cache_id] = {'cards': cards}
        await msg.reply(f'来一把紧张刺激的 24 点！输入算式进行推导，输入「24退出」结束游戏')
        await msg.reply(f'现在你手上有：{cards}，怎么凑 24 点呢？')
    else:
        await msg.reply(f'24点游戏还没结束哦~')


@bot.command(regex=r'(?:24退出)')
async def twenty_four_exit(msg: Message):
    cache_id = f'{msg.ctx.guild.id}-{msg.ctx.channel.id}-{msg.author_id}'
    if cache_id not in cache:
        await msg.reply(f'没有正在进行的24点游戏')
    else:
        del cache[cache_id]
        await msg.reply(f'24点游戏已退出, 这不再来一把？')


@bot.command(regex=r'[\d\+\-\\\*\/]+')
async def twenty_four_step(msg: Message):
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
        await msg.reply('你赢啦！')
        del cache[cache_id]
    elif len(cards) == 1 and cards[0] != 24:
        await msg.reply('你输啦！')
        del cache[cache_id]
    else:
        await msg.reply(f'现在你手上有：{cards}，怎么凑 24 点呢？')


if __name__ == '__main__':
    bot.run()
