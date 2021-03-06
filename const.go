package main

const HELP=`
## 玩法介绍
### 模式
- **Classic**: 经典版斗地主模式
- **LaiZi**: 癞子版斗地主模式
- **Skill**: 癞子版技能大招模式

### 规则
游戏人数2~6人不等,超过3人2副牌,超过5人3副牌,规则参考欢乐斗地主。

出牌时,直接输入想出的牌型,例如3~A顺子:34567890jqka,单10:0, 对2:22,王炸:sx。

癞子模式下同样,缺失的牌会自动使用癞子牌代替,例如当前牌型是*7 6 6 5,输入6665时会自动使用癞子牌*7来代替缺失的6。

更多例子:
- 4个10:0000
- 王炸:sx
- 3~A顺子:34567890jqka
- 3带1:3334
- 飞机:jjjqqq3
`
