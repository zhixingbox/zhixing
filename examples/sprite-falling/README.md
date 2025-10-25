# 精灵从天而降示例 (Sprite Falling Example)

本示例展示了如何使用 SPX API 实现精灵从天而降的效果。

## 功能说明

### 1. 基础下落 (自动开始)
- 游戏开始时，精灵从上方位置 (y=200) 开始下落
- 使用循环逐步改变 Y 坐标实现下落效果
- 使用 `changeYpos -4` 和 `wait 0.05` 控制下落速度

### 2. 物理引擎下落
- 按 **空格键**: 使用物理引擎实现重力下落
- 启用 `DynamicPhysics` 物理模式
- 设置重力值 `setGravity 2`，精灵会受重力影响自然下落

### 3. 滑行下落
- 按 **G键**: 使用 `glide` 平滑下落到目标位置
- 在指定时间内(2秒)滑行到目标坐标

### 4. 速度下落
- 按 **V键**: 设置初始速度，让精灵带着速度下落
- 使用 `setVelocity` 设置向下的速度 (0, -100)
- 结合物理模式实现更真实的下落效果

### 5. 重置位置
- 按 **R键**: 重置精灵到初始位置
- 关闭物理模式，清除速度

## 使用的 API

- `setXYpos x, y`: 设置精灵位置
- `changeYpos delta`: 改变 Y 坐标
- `setPhysicsMode mode`: 设置物理模式 (NoPhysics/DynamicPhysics)
- `setGravity value`: 设置重力值
- `setVelocity vx, vy`: 设置速度向量
- `glide x, y, duration`: 在指定时间内滑行到目标位置
- `repeat count, => {}`: 重复执行指定次数
- `wait seconds`: 等待指定秒数

## 运行方式

将此示例加载到 SPX 游戏引擎中即可运行。按不同按键可以体验不同的下落效果。
