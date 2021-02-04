package document

import (
	"github.com/coadler/helm-docs/pkg/helm"
)

type valueRow struct {
	Key         string
	Type        string
	Description string
	Default     string
}

type chartTemplateData struct {
	helm.ChartDocumentationInfo
	Values []valueRow
}

func getChartTemplateData(chartDocumentationInfo helm.ChartDocumentationInfo) (chartTemplateData, error) {
	valuesTableRows, err := createValueRowsFromObject(
		"",
		chartDocumentationInfo.ChartValues,
		chartDocumentationInfo.ChartValuesDescriptions,
		true,
	)

	if err != nil {
		return chartTemplateData{}, err
	}

	return chartTemplateData{
		ChartDocumentationInfo: chartDocumentationInfo,
		Values:                 valuesTableRows,
	}, nil
}
