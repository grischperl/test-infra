// Code generated by rendertemplates. DO NOT EDIT. 

package releases

// List of currently supported releases
var (
	Release110 = mustParse("1.10")
	Release19 = mustParse("1.9")
	Release18 = mustParse("1.8")
	Release17 = mustParse("1.7")
)

// GetAllKymaReleaseBranches returns all supported kyma release branches
func GetAllKymaReleases() []*SupportedRelease {
	return []*SupportedRelease{
		Release19,
		Release18,
		Release17,
	}
}

