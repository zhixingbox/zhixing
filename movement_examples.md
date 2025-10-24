# xgo/spx 角色位移代码示例

本文档提供了使用 xgo/spx 框架实现角色位移的多种方式。

## 1. 基础位移方法

### 1.1 直接设置坐标位置
```go
setXYpos(100, 100)  // 直接将角色移动到坐标 (100, 100)
setXpos(100)        // 只设置 X 坐标
setYpos(100)        // 只设置 Y 坐标
```

### 1.2 相对位移
```go
changeXYpos(10, 10)  // X 和 Y 方向各移动 10 单位
changeXpos(10)       // X 方向移动 10 单位
changeYpos(10)       // Y 方向移动 10 单位
```

## 2. 滑动移动

### 2.1 滑动到指定位置
```go
glide(100, 100, 1)      // 在 1 秒内滑动到坐标 (100, 100)
glide("S1", 1)          // 在 1 秒内滑动到名为 "S1" 的精灵位置
glide(Mouse, 1)         // 在 1 秒内滑动到鼠标位置
```

## 3. 步进移动

### 3.1 朝当前方向步进
```go
step(100)              // 朝当前朝向移动 100 单位
step(100, 1)           // 朝当前朝向移动 100 单位，速度为 1
step(100, 1, "walk")   // 朝当前朝向移动 100 单位，速度为 1，播放 "walk" 动画
```

### 3.2 步进到目标位置
```go
stepTo("S1")              // 步进到名为 "S1" 的精灵
stepTo(100, 100)          // 步进到坐标 (100, 100)
stepTo(Mouse)             // 步进到鼠标位置
stepTo("S1", 1)           // 步进到名为 "S1" 的精灵，速度为 1
stepTo(100, 100, 1)       // 步进到坐标 (100, 100)，速度为 1
stepTo(100, 100, 1, "walk")  // 步进到坐标 (100, 100)，速度为 1，播放 "walk" 动画
```

## 4. 方向控制

### 4.1 设置朝向
```go
setHeading(Right)      // 设置朝向为右 (90°)
setHeading(Left)       // 设置朝向为左 (-90°)
setHeading(Up)         // 设置朝向为上 (0°)
setHeading(Down)       // 设置朝向为下 (180°)
```

### 4.2 改变朝向
```go
changeHeading(90)      // 将当前朝向增加 90°
```

### 4.3 转向
```go
turn(Right)            // 向右转
turn(Right, 1)         // 向右转，速度为 1
turn(Right, 1, "turn") // 向右转，速度为 1，播放 "turn" 动画
turnTo("S1")           // 转向名为 "S1" 的精灵
turnTo("S1", 1)        // 转向名为 "S1" 的精灵，速度为 1
```

## 5. 物理移动

### 5.1 设置速度
```go
setVelocity(10, 0)     // 设置速度为 X 方向 10 单位/秒，Y 方向 0 单位/秒
```

### 5.2 添加冲量
```go
addImpulse(0, 100)     // 在 Y 方向添加 100 单位的冲量（瞬间速度变化）
```

### 5.3 设置重力
```go
setGravity(2)          // 设置重力为全局重力的 2 倍
setGravity(0)          // 关闭重力
```

### 5.4 设置物理模式
```go
setPhysicsMode(DynamicPhysics)    // 动态物理效果
setPhysicsMode(KinematicPhysics)  // 运动学物理效果
setPhysicsMode(StaticPhysics)     // 静态物理效果
setPhysicsMode(NoPhysics)         // 无物理效果
```

## 6. 键盘控制移动示例

### 6.1 简单方向键控制
```go
type Player struct {
	spx.Sprite
}

func (p *Player) OnKey_KeyUp() {
	p.changeYpos(10)
}

func (p *Player) OnKey_KeyDown() {
	p.changeYpos(-10)
}

func (p *Player) OnKey_KeyLeft() {
	p.changeXpos(-10)
	p.setHeading(Left)
}

func (p *Player) OnKey_KeyRight() {
	p.changeXpos(10)
	p.setHeading(Right)
}
```

### 6.2 WASD 控制
```go
func (p *Player) OnKey_KeyW() {
	p.changeYpos(10)
}

func (p *Player) OnKey_KeyS() {
	p.changeYpos(-10)
}

func (p *Player) OnKey_KeyA() {
	p.changeXpos(-10)
	p.setHeading(Left)
}

func (p *Player) OnKey_KeyD() {
	p.changeXpos(10)
	p.setHeading(Right)
}
```

### 6.3 平滑移动控制
```go
func (p *Player) OnStart() {
	p.forever(func() {
		if p.keyPressed(KeyW) {
			p.step(5)
		}
		if p.keyPressed(KeyS) {
			p.changeHeading(180)
			p.step(5)
			p.changeHeading(-180)
		}
		if p.keyPressed(KeyA) {
			p.turn(Left, 1)
		}
		if p.keyPressed(KeyD) {
			p.turn(Right, 1)
		}
	})
}
```

## 7. 鼠标控制移动

### 7.1 点击移动
```go
func (g *Game) OnClick() {
	player.glide(g.mouseX, g.mouseY, 1)
}
```

### 7.2 跟随鼠标
```go
func (p *Player) OnStart() {
	p.forever(func() {
		p.stepTo(Mouse, 5)
	})
}
```

## 8. 边界检测

### 8.1 碰撞边缘反弹
```go
func (p *Player) OnStart() {
	p.forever(func() {
		p.step(10)
		p.bounceOffEdge()
	})
}
```

### 8.2 检测触碰边缘
```go
if p.touching(Edge) {
	// 触碰到边缘
}

if p.touching(EdgeLeft) {
	// 触碰到左边缘
}
```

## 9. 相机跟随

```go
Camera.follow(player)         // 相机跟随玩家
Camera.follow("PlayerSprite") // 相机跟随名为 "PlayerSprite" 的精灵
Camera.setZoom(2)             // 设置相机缩放为 2 倍
```

## 10. 完整示例

以下是一个完整的可运行示例：

```go
package main

import (
	"github.com/goplus/spx"
)

type Player struct {
	spx.Sprite
	speed float64
}

func (p *Player) OnStart() {
	p.speed = 5
	p.setXYpos(0, 0)
	p.show()
	
	p.forever(func() {
		// WASD 移动控制
		if p.keyPressed(spx.KeyW) {
			p.changeYpos(p.speed)
		}
		if p.keyPressed(spx.KeyS) {
			p.changeYpos(-p.speed)
		}
		if p.keyPressed(spx.KeyA) {
			p.changeXpos(-p.speed)
			p.setHeading(spx.Left)
		}
		if p.keyPressed(spx.KeyD) {
			p.changeXpos(p.speed)
			p.setHeading(spx.Right)
		}
		
		// 空格键跳跃（需要物理系统）
		if p.keyPressed(spx.KeySpace) && p.isOnFloor() {
			p.addImpulse(0, 200)
		}
		
		// 边界检测
		p.bounceOffEdge()
	})
}

type Game struct {
	spx.Game
	player *Player
}

func (g *Game) OnStart() {
	// 设置相机跟随玩家
	spx.Camera.follow(g.player)
}
```

## 总结

xgo/spx 提供了多种灵活的角色位移方式：

1. **基础移动**：直接设置坐标或相对移动
2. **平滑移动**：滑动和步进移动
3. **方向控制**：设置和改变朝向
4. **物理移动**：速度、冲量、重力控制
5. **交互控制**：键盘、鼠标控制
6. **边界处理**：碰撞检测和反弹
7. **相机系统**：相机跟随和缩放

根据不同的游戏需求，选择合适的移动方式可以获得最佳的游戏体验。
