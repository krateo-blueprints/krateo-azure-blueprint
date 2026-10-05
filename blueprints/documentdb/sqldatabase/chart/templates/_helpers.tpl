{{/*
Common labels applied to the Azure Service Operator SqlDatabase resource.
*/}}
{{- define "azure-documentdb-sqldatabase.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
