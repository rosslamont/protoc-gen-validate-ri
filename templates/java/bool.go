package java

const boolTpl = `{{ $f := .Field }}{{ $r := .Rules -}}
{{- if $r.Const }}
	valctx.getValidationCollector().assertValid( ({{ safeName . "value"}}) -> {
		io.envoyproxy.pgv.ConstantValidation.constant("{{ $f.FullyQualifiedName }}", {{ accessor . }}, {{ $r.GetConst }});
	},proto);
{{- end }}`
