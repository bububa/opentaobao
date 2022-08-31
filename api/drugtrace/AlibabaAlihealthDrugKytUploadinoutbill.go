package drugtrace

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytUploadinoutbill 企业上传出入库信息
// alibaba.alihealth.drug.kyt.uploadinoutbill
//
// 企业上传出入库信息，包括101, &#34;生产入库&#34;；102, &#34;采购入库&#34;；103, &#34;退货入库&#34;；104, &#34;调拨入库&#34;；106, &#34;零头入库&#34;；107, &#34;供应入库&#34;；108, &#34;召回入库&#34;；110,&#34;赠品入库&#34;；111,&#34;盘盈入库&#34;；112,&#34;报废入库&#34;；113,&#34;其他入库&#34;
// 201, &#34;销售出库&#34;；202, &#34;退货出库&#34;；203, &#34;调拨出库&#34;；204, &#34;返工出库&#34;；205, &#34;销毁出库&#34;；206, &#34;抽检出库&#34;；207, &#34;直调出库&#34;；208, &#34;生产出库&#34;；209, &#34;供应出库&#34;；211, &#34;召回出库&#34;；212,&#34;赠品出库&#34;；214,&#34;盘亏出库&#34;；215,&#34;损坏出库&#34;；216,&#34;报废出库&#34;；217,&#34;其他出库&#34;；237, &#34;直调退货&#34;。
// 不包括对个人的零售出库，疫苗接种，领药出库。
// 本接口与uploadcircubill接口的主要区别的，本接口入参中直接上传追溯码（多个码时用逗号分隔）。uploadcircubill接口入参中，需要上传码的单据文件（用扫码枪生成的xml文件），一般情况下使用uploadcircubill接口上传单据文件。
func AlibabaAlihealthDrugKytUploadinoutbill(clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytUploadinoutbillAPIRequest, session string) (*drugtrace.AlibabaAlihealthDrugKytUploadinoutbillAPIResponse, error) {
	var resp drugtrace.AlibabaAlihealthDrugKytUploadinoutbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
