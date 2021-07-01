package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSolutionSchemaProductInstancePostAPIRequest
aliexpress.solution.schema.product.instance.post API请求
aliexpress.solution.schema.product.instance.post

Upload product based on json schema instance.QPS(Invoke per second) for this API is limited to 100 for each appkey and 50 for each seller. */
type AliexpressSolutionSchemaProductInstancePostAPIRequest struct {
	model.Params
	// Product instance data. Please note: the shipping_template_id should be replaced with your own shipping template id, which could be obtained through  https://developers.aliexpress.com/en/doc.htm?docId=43456&docType=2
	_productInstanceRequest string
}

// New
