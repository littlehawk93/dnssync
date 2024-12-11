package namesilo

type dnsListResourceRecord struct {
	RecordID string `json:"record_id"`
	Type     string `json:"type"`
	Host     string `json:"host"`
	Value    string `json:"value"`
	TTL      int    `json:"ttl"`
	Distance int    `json:"distance"`
}

type dnsListReply struct {
	Code            int                     `json:"code"`
	Detail          string                  `json:"detail"`
	ResourceRecords []dnsListResourceRecord `json:"resource_record"`
}

type dnsListRequest struct {
	Operation string `json:"operation"`
	IP        string `json:"ip"`
}

type dnsListResponse struct {
	Request dnsListRequest `json:"request"`
	Reply   dnsListReply   `json:"reply"`
}
