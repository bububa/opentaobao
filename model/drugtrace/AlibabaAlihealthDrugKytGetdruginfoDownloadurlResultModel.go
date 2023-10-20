package drugtrace

import (
	"sync"
)

// AlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel 结构体
type AlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel struct {
	// 接口调用失败具体信息
	MsgInfo string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
	// 接口调用失败具体code
	MsgCode string `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
	// 返回值
	Model *JSONObject `json:"model,omitempty" xml:"model,omitempty"`
	// true：接口调用成功
	Success bool `json:"success,omitempty" xml:"success,omitempty"`
}

var poolAlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel)
	},
}

// GetAlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel() 从对象池中获取AlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel
func GetAlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel() *AlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel {
	return poolAlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel.Get().(*AlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel)
}

// ReleaseAlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel 释放AlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel
func ReleaseAlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel(v *AlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel) {
	v.MsgInfo = ""
	v.MsgCode = ""
	v.Model = nil
	v.Success = false
	poolAlibabaAlihealthDrugKytGetdruginfoDownloadurlResultModel.Put(v)
}
