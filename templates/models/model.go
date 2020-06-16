type {{ .CamelizedModelName }} struct {
	{{ range .Properties }} a {{ else }} b {{ end }}
}