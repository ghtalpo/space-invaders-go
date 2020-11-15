package systems

import (
	gc "github.com/x-hgg-x/space-invaders-go/lib/components"

	ec "github.com/ghtalpo/goecsengine/components"
	w "github.com/ghtalpo/goecsengine/world"
	ecs "github.com/x-hgg-x/goecs/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

// MoveBulletSystem moves bullet
func MoveBulletSystem(world w.World) {
	gameComponents := world.Components.Game.(*gc.Components)

	world.Manager.Join(gameComponents.Bullet, world.Components.Engine.Transform).Visit(ecs.Visit(func(entity ecs.Entity) {
		bulletVelocity := gameComponents.Bullet.Get(entity).(*gc.Bullet).Velocity
		bulletTransform := world.Components.Engine.Transform.Get(entity).(*ec.Transform)
		bulletTransform.Translation.Y += bulletVelocity / ebiten.DefaultTPS
	}))
}
