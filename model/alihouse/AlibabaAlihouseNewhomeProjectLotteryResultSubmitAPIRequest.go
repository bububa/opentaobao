package alihouse

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest 楼盘摇号结果提交 API请求
// alibaba.alihouse.newhome.project.lottery.result.submit
//
// 楼盘摇号结果提交
type AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest struct {
	model.Params
	// 摇号结果入参
	_projectLotteryResultParam *ProjectLotteryResultParam
}

// NewAlibabaAlihouseNewhomeProjectLotteryResultSubmitRequest 初始化AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest对象
func NewAlibabaAlihouseNewhomeProjectLotteryResultSubmitRequest() *AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest {
	return &AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest) Reset() {
	r._projectLotteryResultParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest) GetApiMethodName() string {
	return "alibaba.alihouse.newhome.project.lottery.result.submit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetProjectLotteryResultParam is ProjectLotteryResultParam Setter
// 摇号结果入参
func (r *AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest) SetProjectLotteryResultParam(_projectLotteryResultParam *ProjectLotteryResultParam) error {
	r._projectLotteryResultParam = _projectLotteryResultParam
	r.Set("project_lottery_result_param", _projectLotteryResultParam)
	return nil
}

// GetProjectLotteryResultParam ProjectLotteryResultParam Getter
func (r AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest) GetProjectLotteryResultParam() *ProjectLotteryResultParam {
	return r._projectLotteryResultParam
}

var poolAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihouseNewhomeProjectLotteryResultSubmitRequest()
	},
}

// GetAlibabaAlihouseNewhomeProjectLotteryResultSubmitRequest 从 sync.Pool 获取 AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest
func GetAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest() *AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest {
	return poolAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest.Get().(*AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest)
}

// ReleaseAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest 将 AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest(v *AlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest) {
	v.Reset()
	poolAlibabaAlihouseNewhomeProjectLotteryResultSubmitAPIRequest.Put(v)
}
