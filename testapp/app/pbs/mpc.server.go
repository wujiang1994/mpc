package pbs

import (
	"context"
	"google.golang.org/grpc"
)

var Pet *_pet

type _pet struct {
}

func (*_pet) Register(gs *grpc.Server) {
	RegisterPetServer(gs, Pet)
}

func (*_pet) GetPet(ctx context.Context, input *GetPetInput) (output *GetPetOutput, err error) {
	output = new(GetPetOutput)
	output.Base = &BaseOutput{
		Code:                 0,
		Message:              "success",
	}
	output.Animal = map[int64]AnimalType{
		1: AnimalType_Cat,
	}
	return
}
func (*_pet) SetPet(ctx context.Context, input *SetPetInput) (output *SetPetOutput, err error) {
	return &SetPetOutput{}, nil
}
