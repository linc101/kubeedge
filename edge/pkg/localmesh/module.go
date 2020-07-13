package localmesh
import(
	"github.com/kubeedge/beehive/pkg/core"
	// "github.com/kubeedge/kubeedge/edge/pkg/metamanager"
	"github.com/kubeedge/kubeedge/edge/pkg/metamanager/client"
	"github.com/kubeedge/kubeedge/pkg/apis/componentconfig/edgecore/v1alpha1"
	"github.com/kubeedge/kubeedge/edge/pkg/common/modules"
	"github.com/kubeedge/kubeedge/edge/pkg/localmesh/dns"
	
)
type LocalMesh struct {
	enable bool
	metaClient client.CoreInterface

}
const localMeshName = "localmesh"

// Register register localmesh
func Register(m *v1alpha1.LocalMesh) {
	core.Register(newLocalMesh(m))
}

// Name returns the name of EdgeMesh module
func (lm *LocalMesh) Name() string {
	return localMeshName
}

// Group returns EdgeMesh group
func (lm *LocalMesh) Group() string {
	return modules.LocalMeshGroup
}

// Enable indicates whether this module is enabled
func (lm *LocalMesh) Enable() bool {
	return lm.enable
}

//Start sets context and starts the controller
func (lm *LocalMesh) Start() {
	go dns.Start()
}

// 
func newLocalMesh(m *v1alpha1.LocalMesh) *LocalMesh {
	var metaClient client.CoreInterface
	// init meta client
	metaClient = client.New()
	return &LocalMesh{
		enable: m.Enable,
		metaClient: metaClient,
	}

}