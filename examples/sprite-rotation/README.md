# 精灵旋转示例 (Sprite Rotation Example)

本示例展示了如何使用 SPX API 实现精灵的旋转效果。

## 功能说明

### 1. 自动旋转
- 游戏开始时，精灵会自动开始旋转
- 使用 `turn Right, 5` 实现顺时针旋转
- 旋转速度由等待时间控制 (`wait 0.05`)

### 2. 按键控制旋转
- 按 **空格键**: 触发连续旋转，使用 `changeHeading` 改变朝向角度
- 按 **R键**: 演示精灵朝向四个方向的旋转 (上、右、下、左)

## 使用的 API

- `turn Right, speed`: 向指定方向旋转，可指定速度
- `changeHeading degrees`: 改变精灵朝向指定角度
- `setHeading direction`: 设置精灵朝向特定方向 (Up/Right/Down/Left)
- `wait seconds`: 等待指定秒数
- `forever => {}`: 永久循环执行

## 运行方式

将此示例加载到 SPX 游戏引擎中即可运行。
