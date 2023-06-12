package ddk

import "github.com/xiangwork/pinduoduo-sdk/common"

type DuoduoKe struct {
	*common.Context
}

func NewDuoduoKe(context *common.Context) *DuoduoKe {
	service := new(DuoduoKe)
	service.Context = context
	return service
}
