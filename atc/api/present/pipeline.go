package present

import (
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
)

func Pipeline(savedPipeline db.Pipeline) atc.Pipeline {
	return atc.Pipeline{
		ID:          savedPipeline.ID(),
		Name:        savedPipeline.Name(),
		TeamName:    savedPipeline.TeamName(),
		Paused:      savedPipeline.Paused(),
		Public:      savedPipeline.Public(),
		Archived:    savedPipeline.Archived(),
		Groups:      savedPipeline.Groups(),
		Display:     savedPipeline.Display(),
		LastUpdated: savedPipeline.LastUpdated().Unix(),
	}
}

func PipelineKubernetes(savedPipeline db.Pipeline) atc.KubernetesPipeline {
	return atc.KubernetesPipeline{
		TypeMeta: atc.TypeMeta{
			Kind:       "Pipeline",
			APIVersion: "v1alpha1",
		},
		ObjectMeta: atc.ObjectMeta{
			Name:              savedPipeline.Name(),
			Namespace:         savedPipeline.TeamName(),
			CreationTimestamp: savedPipeline.LastUpdated().Unix(),
		},
		Spec: atc.PipelineSpec{
			ID:       savedPipeline.ID(),
			Name:     savedPipeline.Name(),
			TeamName: savedPipeline.TeamName(),
		},
		Status: atc.PipelineStatus{
			Paused:      savedPipeline.Paused(),
			Public:      savedPipeline.Public(),
			Archived:    savedPipeline.Archived(),
			Groups:      savedPipeline.Groups(),
			Display:     savedPipeline.Display(),
			LastUpdated: savedPipeline.LastUpdated().Unix(),
		},
	}

}
