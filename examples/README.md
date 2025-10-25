# SPX API 示例 (SPX API Examples)

本目录包含使用 SPX/xBuilder API 的游戏示例代码。

## 示例列表

### 1. 精灵旋转 (sprite-rotation)
展示如何使用 SPX API 实现精灵的旋转效果，包括：
- 连续自动旋转
- 按键控制旋转
- 朝向不同方向旋转

**主要使用的 API:**
- `turn` - 旋转精灵
- `changeHeading` - 改变朝向角度
- `setHeading` - 设置朝向方向

[查看详细说明 →](./sprite-rotation/README.md)

### 2. 精灵从天而降 (sprite-falling)
展示如何使用 SPX API 实现精灵从天而降的效果，包括：
- 基础逐步下落
- 物理引擎重力下落
- 平滑滑行下落
- 带速度下落

**主要使用的 API:**
- `changeYpos` - 改变 Y 坐标
- `setPhysicsMode` - 设置物理模式
- `setGravity` - 设置重力
- `setVelocity` - 设置速度
- `glide` - 滑行到目标位置

[查看详细说明 →](./sprite-falling/README.md)

## 如何使用

1. 选择一个示例目录
2. 查看目录中的 README.md 了解功能说明
3. 将 main.spx 文件加载到 SPX 游戏引擎中运行

## 相关文档

- [SPX API 完整文档](https://github.com/goplus/spx)
- [xBuilder 官方网站](https://xbuilder.cn)
