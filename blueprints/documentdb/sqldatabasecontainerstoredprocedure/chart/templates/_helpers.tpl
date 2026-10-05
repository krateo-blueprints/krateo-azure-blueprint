{{/*
Common labels applied to the Azure Service Operator SqlDatabaseContainerStoredProcedure resource.
*/}}
{{- define "azure-documentdb-sqldatabasecontainerstoredprocedure.labels" -}}
app.kubernetes.io/managed-by: {{ .Release.Service }}
app.kubernetes.io/part-of: krateo-azure-blueprint
krateo.io/composition: {{ .Release.Name }}
{{- end -}}
