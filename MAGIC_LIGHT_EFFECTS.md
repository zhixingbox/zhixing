# 魔法光效资源使用指南

本文档介绍如何在《小红帽与灰太狼：森林守护者》项目中使用魔法光效资源。

## 📦 资源列表

已创建6种不同的魔法光效SVG文件，位于 `/workspace/assets/sprites/magicLight/` 目录：

### 1. 净化绿光 (purification_green.svg)
- **用途**: 小红帽净化魔法、自然魔法
- **颜色**: 翠绿色、青绿色渐变
- **效果**: 柔和的光晕，带有闪光粒子
- **适用场景**: 
  - 场景3-15：净化时刻
  - 守护者魔法施放
  - 自然力量展示

### 2. 黑暗紫光 (dark_purple.svg)
- **用途**: 黑暗巫师魔法、控制效果
- **颜色**: 深紫色、暗紫色渐变
- **效果**: 不祥的光晕，带有触手状能量线
- **适用场景**:
  - 灰太狼被控制时
  - 黑暗巫师施法
  - 邪恶力量展示

### 3. 金色力量 (golden_power.svg)
- **用途**: 守护者觉醒、终极力量
- **颜色**: 金黄色、橙色渐变
- **效果**: 爆发式光芒，八方射线
- **适用场景**:
  - 场景4-7：小红帽力量觉醒
  - 守护者试炼通过
  - 强大魔法爆发

### 4. 蓝色护盾 (blue_shield.svg)
- **用途**: 防御魔法、护盾展示
- **颜色**: 天蓝色、钴蓝色渐变
- **效果**: 盾形防护罩，带有能量波纹
- **适用场景**:
  - 场景1-3：第一次战斗防御
  - 魔法防护展示
  - 守护结界

### 5. 红色警告 (red_warning.svg)
- **用途**: 危险提示、被控制状态
- **颜色**: 鲜红色、深红色渐变
- **效果**: 脉冲式警告光，带有闪电效果
- **适用场景**:
  - 灰太狼眼睛发红光
  - 危险警告
  - 战斗紧张时刻

### 6. 治疗闪光 (healing_sparkle.svg)
- **用途**: 治疗效果、恢复能量
- **颜色**: 白色、淡金色渐变
- **效果**: 十字光芒，柔和的治愈感
- **适用场景**:
  - 魔法露水使用时
  - 净化草药效果
  - 能量恢复

## 🎮 使用方法

### 基础用法

在 `.spx` 文件中引用 magicLight 精灵：

```go
// 显示净化光效
magicLight.setXYpos 0, -20
magicLight.setCostume "净化绿光"
magicLight.show

// 添加闪烁效果
repeat 10, => {
    magicLight.changeGraphicEffect BrightnessEffect, 10
    wait 0.1
}

// 隐藏光效
magicLight.hide
```

### 场景1-3：防御魔法示例

```go
// 小红帽使用防御魔法
onMsg "战斗开始", => {
    // 小红帽施法动作
    redHood.setCostume "redhood_casting"
    
    // 显示蓝色护盾光效
    magicLight.setXYpos redHood.xpos, redHood.ypos
    magicLight.setCostume "蓝色护盾"
    magicLight.setSize 1.5
    magicLight.show
    
    // 护盾闪烁效果
    repeat 5, => {
        magicLight.setGraphicEffect BrightnessEffect, 30
        wait 0.2
        magicLight.clearGraphicEffects
        wait 0.2
    }
    
    play "magic_shield_sound"
}
```

### 场景3-15：净化魔法完整示例

```go
// 净化魔法序列
onMsg "净化开始", => {
    // 1. 小红帽凝聚力量
    redHood.setCostume "redhood_casting"
    
    // 显示绿色光芒在小红帽手中
    magicLight.setXYpos redHood.xpos, redHood.ypos - 20
    magicLight.setCostume "净化绿光"
    magicLight.setSize 0.5
    magicLight.show
    
    // 光芒逐渐增强
    repeat 10, => {
        magicLight.changeSize 0.1
        magicLight.changeGraphicEffect BrightnessEffect, 5
        wait 0.15
    }
    
    wait 1
    
    // 2. 光芒射向灰太狼
    magicLight.glide wolf.xpos, wolf.ypos, 1
    
    // 3. 笼罩灰太狼
    magicLight.setSize 3
    repeat 8, => {
        magicLight.setGraphicEffect ColorEffect, rand(80, 120)
        wait 0.3
    }
    
    // 4. 净化完成
    magicLight.clearGraphicEffects
    wait 0.5
    magicLight.hide
}
```

### 场景4-7：力量觉醒示例

```go
// 小红帽觉醒守护者力量
onMsg "力量觉醒", => {
    // 觉醒光效
    magicLight.setXYpos redHood.xpos, redHood.ypos
    magicLight.setCostume "金色力量"
    magicLight.setSize 0.5
    magicLight.show
    
    // 爆发式增长
    repeat 20, => {
        magicLight.changeSize 0.2
        magicLight.changeGraphicEffect BrightnessEffect, 10
        wait 0.08
    }
    
    // 光芒爆炸
    repeat 5, => {
        magicLight.changeSize 0.5
        wait 0.05
    }
    
    play "power_awakening_sound"
    
    // 光效消散
    repeat 15, => {
        magicLight.changeGraphicEffect GhostEffect, 7
        wait 0.1
    }
    
    magicLight.hide
}
```

### 灰太狼被控制效果示例

```go
// 显示灰太狼被控制的红光
onBackdrop "moonlight_lake", => {
    // 设置角色位置
    wolf.setXYpos 160, -20
    wolf.show
    
    // 在灰太狼位置显示红色警告光
    magicLight.setXYpos wolf.xpos, wolf.ypos
    magicLight.setCostume "红色警告"
    magicLight.setSize 1.2
    magicLight.show
    
    // 持续脉冲效果
    forever => {
        repeat 3, => {
            magicLight.setGraphicEffect ColorEffect, 100
            wait 0.15
            magicLight.clearGraphicEffects
            wait 0.15
        }
        wait 1
    }
}
```

### 治疗草药效果示例

```go
// 使用净化草药
onMsg "使用草药", => {
    // 草药飞向灰太狼
    cureHerb.glide wolf.xpos, wolf.ypos, 1.5
    
    // 显示治疗光效
    magicLight.setXYpos wolf.xpos, wolf.ypos
    magicLight.setCostume "治疗闪光"
    magicLight.setSize 1.5
    magicLight.show
    
    // 柔和的治疗闪光
    repeat 15, => {
        magicLight.changeGraphicEffect BrightnessEffect, 6
        wait 0.12
    }
    
    repeat 15, => {
        magicLight.changeGraphicEffect BrightnessEffect, -6
        wait 0.12
    }
    
    // 草药消失
    cureHerb.hide
    
    // 光效淡出
    repeat 10, => {
        magicLight.changeGraphicEffect GhostEffect, 10
        wait 0.1
    }
    
    magicLight.hide
    magicLight.clearGraphicEffects
}
```

## 🎨 高级技巧

### 1. 多重光效叠加
```go
// 同时使用多个光效（需要克隆或使用多个精灵实例）
magicLight.setCostume "净化绿光"
magicLight.setXYpos 0, 0
magicLight.show

// 在需要时切换不同光效
wait 2
magicLight.setCostume "金色力量"
```

### 2. 光效跟随角色
```go
// 让光效跟随角色移动
forever => {
    magicLight.setXYpos redHood.xpos, redHood.ypos + 10
    wait 0.03
}
```

### 3. 组合动画效果
```go
// 结合多种图形效果
magicLight.show
repeat 20, => {
    magicLight.changeGraphicEffect BrightnessEffect, 5
    magicLight.changeGraphicEffect ColorEffect, 10
    magicLight.changeSize 0.05
    wait 0.1
}
```

### 4. 光效脉冲效果
```go
// 创建脉冲效果
forever => {
    repeat 10, => {
        magicLight.changeSize 0.05
        wait 0.05
    }
    repeat 10, => {
        magicLight.changeSize -0.05
        wait 0.05
    }
}
```

## 📝 注意事项

1. **性能优化**: 不要同时显示过多光效，建议最多3-4个
2. **层级管理**: 使用 `setLayer` 控制光效在其他精灵的上方或下方
3. **光效清理**: 使用完毕后记得调用 `hide` 和 `clearGraphicEffects`
4. **尺寸调整**: 根据场景需要调整 `setSize`，建议范围 0.5-3.0
5. **颜色搭配**: 注意光效颜色与场景背景的对比度

## 🎬 推荐使用场景

### 第一集：森林的异变
- **场景1-3**: 使用"蓝色护盾"展示小红帽的防御魔法

### 第三集：意外的盟友  
- **场景3-15**: 使用"净化绿光"展示净化魔法
- 使用"红色警告"展示灰太狼被控制状态
- 使用"治疗闪光"展示魔法露水效果

### 第四集：最终决战
- **场景4-7**: 使用"金色力量"展示小红帽觉醒
- 使用"黑暗紫光"展示黑暗巫师魔法
- 使用"蓝色护盾"展示联合防御

### 第五集：新的开始
- **场景4-8**: 使用"金色力量"和"治疗闪光"展示森林祝福

## 💡 扩展建议

如需添加更多光效类型，可以：

1. 创建新的SVG文件到 `assets/sprites/magicLight/` 目录
2. 在 `index.json` 的 `costumes` 数组中添加新服装
3. 在 `magicLight.spx` 中添加对应的使用函数

## 📞 技术支持

如有问题或需要定制其他光效，请在 GitHub Issue 中提出。

---

**创建日期**: 2025-10-25
**版本**: 1.0
**适用项目**: 《小红帽与灰太狼：森林守护者》
