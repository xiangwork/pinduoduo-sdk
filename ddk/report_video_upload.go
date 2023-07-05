package ddk

import (
	"os"

	"github.com/xiangwork/pinduoduo-sdk/common"
	"github.com/xiangwork/pinduoduo-sdk/util"
)

type ReportVideoUploadParams struct {
	File os.File `json:"file"`
}

type ReportVideoUploadPostResult struct {
	Response struct {
		Url string `json:"url"` //请求到的结果数
	} `json:"response"`
	common.CommonResult
}

func (d *DuoduoKe) ReportVideoUploadPost(p *ReportVideoUploadParams) (*ReportVideoUploadPostResult, error) {
	apiType := `pdd.ddk.report.video.upload`
	params, paramsURL := util.FormatURLParams(p)
	url := d.GetUploadURL(apiType, "", params, paramsURL)
	var result ReportVideoUploadPostResult
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
