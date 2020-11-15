package systems

import (
	"fmt"

	gc "github.com/x-hgg-x/space-invaders-go/lib/components"
	"github.com/x-hgg-x/space-invaders-go/lib/resources"

	ec "github.com/ghtalpo/goecsengine/components"
	w "github.com/ghtalpo/goecsengine/world"
	ecs "github.com/x-hgg-x/goecs/v2"
)

// LifeSystem manages lives
func LifeSystem(world w.World) {
	gameComponents := world.Components.Game.(*gc.Components)
	gameResources := world.Resources.Game.(*resources.Game)

	for range gameResources.Events.LifeEvents {
		gameResources.Lives--

		world.Manager.Join(world.Components.Engine.Text, world.Components.Engine.UITransform).Visit(ecs.Visit(func(entity ecs.Entity) {
			text := world.Components.Engine.Text.Get(entity).(*ec.Text)
			if text.ID == "game_life" {
				text.Text = fmt.Sprintf("LIVES: %d", gameResources.Lives)
			}
		}))

		world.Manager.Join(gameComponents.Player, gameComponents.Controllable, world.Components.Engine.SpriteRender).Visit(ecs.Visit(func(playerEntity ecs.Entity) {
			playerSprite := world.Components.Engine.SpriteRender.Get(playerEntity).(*ec.SpriteRender)

			playerEntity.AddComponent(world.Components.Engine.AnimationControl, &ec.AnimationControl{
				Animation:      playerSprite.SpriteSheet.Animations[resources.PlayerDeathAnimation],
				Command:        ec.AnimationCommand{Type: ec.AnimationCommandStart},
				RateMultiplier: 1,
			})
		}))

		gameResources.StateEvent = resources.StateEventDeath
		break
	}
	gameResources.Events.LifeEvents = nil
}
