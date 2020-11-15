package systems

import (
	"fmt"

	"github.com/x-hgg-x/space-invaders-go/lib/resources"

	ec "github.com/ghtalpo/goecsengine/components"
	"github.com/ghtalpo/goecsengine/math"
	w "github.com/ghtalpo/goecsengine/world"
	ecs "github.com/x-hgg-x/goecs/v2"
)

// ScoreSystem manages score
func ScoreSystem(world w.World) {
	gameResources := world.Resources.Game.(*resources.Game)

	for _, scoreEvent := range gameResources.Events.ScoreEvents {
		gameResources.Score += scoreEvent.Score
		gameResources.Score = math.Min(99999, gameResources.Score)

		world.Manager.Join(world.Components.Engine.Text, world.Components.Engine.UITransform).Visit(ecs.Visit(func(entity ecs.Entity) {
			text := world.Components.Engine.Text.Get(entity).(*ec.Text)
			if text.ID == "game_score" {
				text.Text = fmt.Sprintf("SCORE: %d", gameResources.Score)
			}
		}))
	}
	gameResources.Events.ScoreEvents = nil
}
