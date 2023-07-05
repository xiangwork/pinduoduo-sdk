package ddk

import (
	"github.com/xiangwork/pinduoduo-sdk/common"
	"github.com/xiangwork/pinduoduo-sdk/util"
)

type ReportImgUploadParams struct {
	File string `json:"file"`
}

type ReportImgUploadPostResult struct {
	Response struct {
		Url string `json:"url"` //请求到的结果数
	} `json:"response"`
	common.CommonResult
}

func (d *DuoduoKe) ReportImgUploadPost(p *ReportImgUploadParams) (*ReportImgUploadPostResult, error) {
	apiType := `pdd.ddk.report.img.upload`
	params, paramsURL := util.FormatURLParams(p)
	url := d.GetUploadURL(apiType, "", params, paramsURL)
	var result ReportImgUploadPostResult
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
