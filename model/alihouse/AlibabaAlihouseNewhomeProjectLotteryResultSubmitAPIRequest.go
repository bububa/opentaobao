package alihouse

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihousenewhomeprojectlotteryresultsubmitAPIRequest 楼盘摇号结果提交 API请求
// alibaba.alihouse.newhome.project.lottery.result.submit
//
// 楼盘摇号结果提交
type AlibabaalihousenewhomeprojectlotteryresultsubmitAPIRequest struct {
	model.Params
	// 摇号结果入参
	_projectLotteryResultParam *ProjectLotteryResultParam
}

// NewAlibabaalihousenewhomeprojectlotteryresultsubmitRequest 初始化AlibabaalihousenewhomeprojectlotteryresultsubmitAPIRequest对象
func NewAlibabaalihousenewhomeprojectlotteryresultsubmitRequest() *AlibabaalihousenewhomeprojectlotteryresultsubmitAPIRequest {
	return &AlibabaalihousenewhomeprojectlotteryresultsubmitAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihousenewhomeprojectlotteryresultsubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.lottery.result.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihousenewhomeprojectlotteryresultsubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihousenewhomeprojectlotteryresultsubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectLotteryResultParam is ProjectLotteryResultParam Setter
// 摇号结果入参
func (r *AlibabaalihousenewhomeprojectlotteryresultsubmitAPIRequest) SetProjectLotteryResultParam(_projectLotteryResultParam *ProjectLotteryResultParam) error {
	r._projectLotteryResultParam = _projectLotteryResultParam
	r.Set("project_lottery_result_param", _projectLotteryResultParam)
	return nil
}

// GetProjectLotteryResultParam ProjectLotteryResultParam Getter
func (r AlibabaalihousenewhomeprojectlotteryresultsubmitAPIRequest) GetProjectLotteryResultParam() *ProjectLotteryResultParam {
	return r._projectLotteryResultParam
}
