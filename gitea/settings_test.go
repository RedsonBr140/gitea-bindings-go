package gitea

import "testing"

func TestAPISettingsMarshal(t *testing.T) {
	testJSONMarshal(t, &GeneralAPISettings{}, "{}")

	u := &GeneralAPISettings{
		MaxResponseItems:       1,
		DefaultPagingNum:       2,
		DefaultGitTreesPerPage: 3,
		DefaultMaxBlobSize:     1998,
	}

	want := `{
	"max_response_items":1,
	"default_paging_num":2,
	"default_git_trees_per_page":3,
	"default_max_blob_size":1998
	}`

	testJSONMarshal(t, u, want)
}
