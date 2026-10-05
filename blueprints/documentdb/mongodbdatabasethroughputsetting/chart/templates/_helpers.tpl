{{/*
Common labels applied to the Azure Service Operator MongodbDatabaseThroughputSetting resource.
*/}}
{{- define "azure-documentdb-mongodbdatabasethroughputsetting.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
