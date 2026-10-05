{{/*
Common labels applied to the Azure Service Operator FlexibleServersDatabase resource.
*/}}
{{- define "azure-dbformysql-flexibleserversdatabase.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
