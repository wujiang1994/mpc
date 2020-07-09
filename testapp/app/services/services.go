package services

import "mpc"

var MpcTest *_mpcTest

type _mpcTest struct {
}

func SetupMpcTest() {
	MpcTest = &_mpcTest{}
}

func (m *_mpcTest) Do(ctx *mpc.Context, input *MpcTestDoReq) (output *MpcTestDoResp, err error) {
	logger := mpc.NewAppLogger()
	output = &MpcTestDoResp{
		Name: input.Name,
		Age:  18,
		Sex:  0,
		Transcript: &Transcript{
			Math: 99,
			Eng:  95,
		},
	}
	logger.Infof("receive request: %+v, response: %+v", input, output)
	return
}

type MpcTestDoReq struct {
	Name string `json:"name"`
}

type MpcTestDoResp struct {
	Name       string      `json:"name"`
	Age        int         `json:"age"`
	Sex        int         `json:"sex"`
	Transcript *Transcript `json:"transcript"`
}

type Transcript struct {
	Math int `json:"math"`
	Eng  int `json:"eng"`
}
