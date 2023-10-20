package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampuscoreemployeemodifyemployee 修改员工基本信息
// alibaba.campus.core.employee.modifyemployee
//
// 根据用户ID和公司ID更新员工基本信息（头像、性别、昵称）
func Alibabacampuscoreemployeemodifyemployee(clt *core.SDKClient, req *campus.AlibabacampuscoreemployeemodifyemployeeAPIRequest, session string) (*campus.AlibabacampuscoreemployeemodifyemployeeAPIResponse, error) {
	var resp campus.AlibabacampuscoreemployeemodifyemployeeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
