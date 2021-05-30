package api

type LimitOffset struct {
	Limit  int64 `json:"limit" form:"limit,default=10"`
	Offset int64 `json:"offset" form:"offset,default=0"`
}

type PaginationParams struct {
	LimitOffset
	Count int64 `json:"count"`
	//Next     *string `json:"next" binding:"null"`
	//Previous *string `json:"previous"`
}

func (p *PaginationParams) GetNextParams() *LimitOffset {
	newOffset := p.Limit + p.Offset
	if p.Count > newOffset {
		return &LimitOffset{Limit: p.Limit, Offset: newOffset}
	}
	return nil
}

func (p *PaginationParams) GetPrevParams() *LimitOffset {
	newOffset := p.Offset - p.Limit
	if newOffset >= 0 {
		return &LimitOffset{Limit: p.Limit, Offset: newOffset}
	}
	return nil
}
