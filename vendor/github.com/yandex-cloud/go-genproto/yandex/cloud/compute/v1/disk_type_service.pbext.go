// Code generated by protoc-gen-goext. DO NOT EDIT.

package compute

func (m *GetDiskTypeRequest) SetDiskTypeId(v string) {
	m.DiskTypeId = v
}

func (m *ListDiskTypesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDiskTypesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDiskTypesResponse) SetDiskTypes(v []*DiskType) {
	m.DiskTypes = v
}

func (m *ListDiskTypesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}