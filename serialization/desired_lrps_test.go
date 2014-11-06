package serialization_test

import (
	"github.com/cloudfoundry-incubator/receptor"
	"github.com/cloudfoundry-incubator/receptor/serialization"
	"github.com/cloudfoundry-incubator/runtime-schema/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DesiredLRP Serialization", func() {
	Describe("DesiredLRPFromRequest", func() {
		var request receptor.DesiredLRPCreateRequest
		var desiredLRP models.DesiredLRP

		BeforeEach(func() {
			request = receptor.DesiredLRPCreateRequest{
				ProcessGuid: "the-process-guid",
				Domain:      "the-domain",
				Stack:       "the-stack",
				RootFSPath:  "the-rootfs-path",
				Annotation:  "foo",
				Instances:   1,
				Actions: []models.ExecutorAction{
					{
						Action: &models.RunAction{
							Path: "the-path",
						},
					},
				},
			}
		})
		JustBeforeEach(func() {
			var err error
			desiredLRP, err = serialization.DesiredLRPFromRequest(request)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("translates the request into a DesiredLRP model, preserving attributes", func() {
			Ω(desiredLRP.ProcessGuid).Should(Equal("the-process-guid"))
			Ω(desiredLRP.Domain).Should(Equal("the-domain"))
			Ω(desiredLRP.Stack).Should(Equal("the-stack"))
			Ω(desiredLRP.RootFSPath).Should(Equal("the-rootfs-path"))
			Ω(desiredLRP.Annotation).Should(Equal("foo"))
		})
	})

	Describe("DesiredLRPToResponse", func() {
		var desiredLRP models.DesiredLRP
		BeforeEach(func() {
			desiredLRP = models.DesiredLRP{
				ProcessGuid: "process-guid-0",
				Domain:      "domain-0",
				RootFSPath:  "root-fs-path-0",
				Instances:   127,
				Stack:       "stack-0",
				EnvironmentVariables: []models.EnvironmentVariable{
					{Name: "ENV_VAR_NAME", Value: "value"},
				},
				Actions: []models.ExecutorAction{
					{Action: models.RunAction{Path: "/bin/true"}},
				},
				DiskMB:    126,
				MemoryMB:  1234,
				CPUWeight: 192,
				Ports: []models.PortMapping{
					{ContainerPort: 456, HostPort: 876},
				},
				Routes:     []string{"route-0", "route-1"},
				Log:        models.LogConfig{"log-guid-0", "log-source-name-0"},
				Annotation: "annotation-0",
			}
		})

		It("serializes all the fields", func() {
			expectedResponse := receptor.DesiredLRPResponse{
				ProcessGuid: "process-guid-0",
				Domain:      "domain-0",
				RootFSPath:  "root-fs-path-0",
				Instances:   127,
				Stack:       "stack-0",
				EnvironmentVariables: []receptor.EnvironmentVariable{
					{Name: "ENV_VAR_NAME", Value: "value"},
				},
				Actions: []models.ExecutorAction{
					{Action: models.RunAction{Path: "/bin/true"}},
				},
				DiskMB:    126,
				MemoryMB:  1234,
				CPUWeight: 192,
				Ports: []receptor.PortMapping{
					{ContainerPort: 456, HostPort: 876},
				},
				Routes:     []string{"route-0", "route-1"},
				Log:        receptor.LogConfig{"log-guid-0", "log-source-name-0"},
				Annotation: "annotation-0",
			}

			actualResponse := serialization.DesiredLRPToResponse(desiredLRP)
			Ω(actualResponse).Should(Equal(expectedResponse))
		})
	})
})