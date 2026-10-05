{{/*
Common labels applied to the Azure Service Operator Route resource.
*/}}
{{- define "azure-cdn-route.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
