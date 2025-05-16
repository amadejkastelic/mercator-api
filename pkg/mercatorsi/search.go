package mercatorsi

import (
	"fmt"
	"net/url"
	"strconv"
)

const searchURL = "/products/browseProducts/getProducts"

func (c *client) Search(in SearchRequest) (*SearchResponse, error) {
	query := url.Values{}
	query.Set("limit", strconv.Itoa(in.Limit))
	query.Set("offset", strconv.Itoa(in.Offset))
	query.Set("from", strconv.Itoa(in.From))

	if in.Filter != "" {
		query.Set("filterData[search]", in.Filter)
	}
	if in.Sort != nil {
		if in.Sort.Field != "" {
			query.Set("filterData[sort]", string(in.Sort.Field))
		}
		if in.Sort.Direction != "" {
			query.Set("filterData[sort-direction]", string(in.Sort.Direction))
		}
	}

	req, err := c.newRequest("GET", searchURL, query, nil)
	if err != nil {
		return nil, err
	}

	var resp *SearchResponse
	status, err := c.do(req, &resp)
	if err != nil {
		return nil, err
	}
	if status.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status: %s", status.Status)
	}

	return resp, nil
}
