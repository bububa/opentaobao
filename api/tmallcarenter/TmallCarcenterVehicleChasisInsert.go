package tmallcarenter

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallcarenter"
)

// Tmallcarcentervehiclechasisinsert EPC车型底盘压缩库新增接口
// tmall.carcenter.vehicle.chasis.insert
//
// EPC车型底盘压缩库新增接口
func Tmallcarcentervehiclechasisinsert(clt *core.SDKClient, req *tmallcarenter.TmallcarcentervehiclechasisinsertAPIRequest, session string) (*tmallcarenter.TmallcarcentervehiclechasisinsertAPIResponse, error) {
	var resp tmallcarenter.TmallcarcentervehiclechasisinsertAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
