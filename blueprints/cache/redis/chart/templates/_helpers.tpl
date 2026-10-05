{{/*
Common labels applied to the Azure Service Operator Redis resource.
*/}}
{{- define "azure-cache-redis.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
