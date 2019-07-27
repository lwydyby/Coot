# -*- coding: utf-8 -*-
# 本脚本支持 python2 和 python 3
# 每天随机三个食物

import random

food_list = ["1", "2", "3", "4", "5", "6", "7"]

set_list = set()

while True:
    index = random.randint(1, len(food_list))

    set_list.add(food_list[index - 1])

    if len(set_list) >= 3:
        break

print('0&&' + ' - '.join(set_list))