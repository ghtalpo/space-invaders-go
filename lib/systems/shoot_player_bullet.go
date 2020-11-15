package systems

import (
	gc "github.com/x-hgg-x/space-invaders-go/lib/components"
	"github.com/x-hgg-x/space-invaders-go/lib/resources"

	ec "github.com/ghtalpo/goecsengine/components"
	"github.com/ghtalpo/goecsengine/loader"
	"github.com/ghtalpo/goecsengine/math"
	w "github.com/ghtalpo/goecsengine/world"
	ecs "github.com/x-hgg-x/goecs/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

var shootPlayerBulletFrame = 0

// ShootPlayerBulletSystem shoots player bullet
func ShootPlayerBulletSystem(world w.World) {
	shootPlayerBulletFrame--

	gameComponents := world.Components.Game.(*gc.Components)
	audioPlayers := *world.Resources.AudioPlayers

	if world.Manager.Join(gameComponents.Player, gameComponents.Bullet).Empty() {
		shootPlayerBulletFrame = math.Min(ebiten.DefaultTPS/20, shootPlayerBulletFrame)
	}

	if world.Resources.InputHandler.Actions[resources.ShootAction] && shootPlayerBulletFrame <= 0 {
		shootPlayerBulletFrame = ebiten.DefaultTPS

		firstPlayer := ecs.GetFirst(world.Manager.Join(gameComponents.Player, gameComponents.Controllable, world.Components.Engine.Transform))
		if firstPlayer == nil {
			return
		}
		playerX := world.Components.Engine.Transform.Get(ecs.Entity(*firstPlayer)).(*ec.Transform).Translation.X

		playerBulletEntity := loader.AddEntities(world, world.Resources.Prefabs.(*resources.Prefabs).Game.PlayerBullet)
		for iEntity := range playerBulletEntity {
			playerBulletTransform := world.Components.Engine.Transform.Get(playerBulletEntity[iEntity]).(*ec.Transform)
			playerBulletTransform.Translation.X = playerX
		}

		audioPlayers["shoot"].Rewind()
		audioPlayers["shoot"].Play()
	}
}
