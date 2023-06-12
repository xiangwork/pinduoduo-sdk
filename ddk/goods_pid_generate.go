package ddk

import (
	"github.com/xiangwork/pinduoduo-sdk/common"
	"github.com/xiangwork/pinduoduo-sdk/util"
)

/*
*
查询已经生成的推广位信息
*/
type GoodsPidGenerateParams struct {
	Number      *int `json:"number"`                   //生成数量
	PIdNameList *int `json:"p_id_name_list,omitempty"` //推广位名称，例如["1","2"]
	MediaId     *int `json:"media_id,omitempty"`       //媒体id
}

type GoodsPidGenerateResult struct {
	PIdGenerateResponse struct {
		PIdList        []GoodsPidQueryInfo `json:"p_id_list"`        //多多进宝推广位对象列表
		RemainPidCount int                 `json:"remain_pid_count"` //PID剩余数量
	} `json:"p_id_generate_response"`
	common.CommonResult
}

type GoodsPidGenerateInfo struct {
	PID        string `json:"p_id"`        //推广位ID
	Name       string `json:"pid_name"`    //推广位名称
	CreateTime int64  `json:"create_time"` //推广位生成时间
	MediaId    int64  `json:"media_id"`    //媒体id
}

func (d *DuoduoKe) GoodsPidGenerate(p *GoodsPidGenerateParams) (*GoodsPidGenerateResult, error) {
	apiType := `pdd.ddk.goods.pid.generate`
	params, paramsURL := util.FormatURLParams(p)
	url := d.GetURL(apiType, "", params, paramsURL)

	var result GoodsPidGenerateResult
	err := util.HttpPOST(url, nil, &result)
	if err != nil {
		return nil, err
	}

	err = common.CheckErrCode(result.CommonResult)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
