package java

const msgTpl = `
{{ if not (ignored .) -}}
	/**
	 * Validates {@code {{ simpleName . }}} protobuf objects.
	 */
	public static class {{ simpleName . }}Validator implements io.envoyproxy.pgv.ValidatorImpl<{{ qualifiedName . }}> {
		{{- template "msgInner" . -}}
	}
{{- end -}}
`

const msgInnerTpl = `
	{{- range .NonOneOfFields }}
		{{ renderConstants (context .) }}
	{{ end }}
	{{ range .RealOneOfs }}
		{{ template "oneOfConst" . }}
	{{ end }}

	public void assertValid({{ qualifiedName . }} proto, io.envoyproxy.pgv.ValidationContext valctx) throws io.envoyproxy.pgv.ValidationException {
	{{ if disabled . }}
		// Validate is disabled for {{ simpleName . }}
		return;
	{{- else -}}
	{{ range .NonOneOfFields -}}
		{{ render (context .) }}
	{{ end -}}
	{{ range .RealOneOfs }}
		{{ template "oneOf" . }}
	{{- end -}}
	{{- end }}
	}
`
