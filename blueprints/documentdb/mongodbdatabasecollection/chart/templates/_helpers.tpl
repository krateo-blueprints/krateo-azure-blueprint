{{/*
Common labels applied to the Azure Service Operator MongodbDatabaseCollection resource.
*/}}
{{- define "azure-documentdb-mongodbdatabasecollection.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
