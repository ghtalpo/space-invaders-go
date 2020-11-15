package systems

import (
	gc "github.com/x-hgg-x/space-invaders-go/lib/components"

	ec "github.com/ghtalpo/goecsengine/components"
	w "github.com/ghtalpo/goecsengine/world"
	ecs "github.com/x-hgg-x/goecs/v2"
)

// DeleteSystem removes deleted entities after animation end
func DeleteSystem(world w.World) {
	gameComponents := world.Components.Game.(*gc.Components)

	world.Manager.Join(gameComponents.Deleted, world.Components.Engine.AnimationControl).Visit(ecs.Visit(func(entity ecs.Entity) {
		animationControl := world.Components.Engine.AnimationControl.Get(entity).(*ec.AnimationControl)
		if animationControl.GetState().Type == ec.ControlStateDone {
			world.Manager.DeleteEntity(entity)
		}
	}))
}
