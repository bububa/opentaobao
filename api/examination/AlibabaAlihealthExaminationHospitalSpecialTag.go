package examination

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/examination"
)

// Alibabaalihealthexaminationhospitalspecialtag 体检机构获取特色服务标签
// alibaba.alihealth.examination.hospital.special.tag
//
// 体检机构获取特色服务标签列表
func Alibabaalihealthexaminationhospitalspecialtag(clt *core.SDKClient, req *examination.AlibabaalihealthexaminationhospitalspecialtagAPIRequest, session string) (*examination.AlibabaalihealthexaminationhospitalspecialtagAPIResponse, error) {
	var resp examination.AlibabaalihealthexaminationhospitalspecialtagAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
