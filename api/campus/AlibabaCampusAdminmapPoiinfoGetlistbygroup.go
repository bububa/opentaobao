package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabacampusadminmappoiinfogetlistbygroup 根据分组条件查询分组下的空间单元不包涵业务属性信息
// alibaba.campus.adminmap.poiinfo.getlistbygroup
//
// 根据分组条件查询分组下的空间单元不包涵业务属性信息
func Alibabacampusadminmappoiinfogetlistbygroup(clt *core.SDKClient, req *campus.AlibabacampusadminmappoiinfogetlistbygroupAPIRequest, session string) (*campus.AlibabacampusadminmappoiinfogetlistbygroupAPIResponse, error) {
	var resp campus.AlibabacampusadminmappoiinfogetlistbygroupAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
