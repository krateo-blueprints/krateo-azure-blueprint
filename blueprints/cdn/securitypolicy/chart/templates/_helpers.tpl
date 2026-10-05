{{/*
Common labels applied to the Azure Service Operator SecurityPolicy resource.
*/}}
{{- define "azure-cdn-securitypolicy.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
