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
	p.setSize(1)
	p.show()
}

func (p *Player) OnKey_KeyW() {
	p.changeYpos(p.speed)
}

func (p *Player) OnKey_KeyS() {
	p.changeYpos(-p.speed)
}

func (p *Player) OnKey_KeyA() {
	p.changeXpos(-p.speed)
	p.setHeading(spx.Left)
}

func (p *Player) OnKey_KeyD() {
	p.changeXpos(p.speed)
	p.setHeading(spx.Right)
}

func (p *Player) OnKey_KeySpace() {
	p.glide(0, 0, 1)
}

func (p *Player) SmoothMovement() {
	p.OnStart(func() {
		p.forever(func() {
			if p.keyPressed(spx.KeyUp) {
				p.step(10)
			}
			if p.keyPressed(spx.KeyDown) {
				p.changeHeading(180)
				p.step(10)
				p.changeHeading(-180)
			}
			if p.keyPressed(spx.KeyLeft) {
				p.turn(spx.Left, 1)
			}
			if p.keyPressed(spx.KeyRight) {
				p.turn(spx.Right, 1)
			}
		})
	})
}

type Game struct {
	spx.Game
	player *Player
}

func (g *Game) OnStart() {
	g.player = &Player{}
	g.player.OnStart()
}
