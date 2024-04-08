package client

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/message"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

const deleteCypher = "MATCH (n:%s) WHERE n._cq_source_name = $cq_source_name AND n._cq_sync_time < $cq_sync_time DETACH DELETE n"
const deleteAccountCypher = "MATCH (n:%s) WHERE n._cq_source_name = $cq_source_name AND n._cq_sync_time < $cq_sync_time AND n.%s = $account_id DETACH DELETE n"

func (c *Client) DeleteStale(ctx context.Context, msgs message.WriteDeleteStales) error {
	session := c.Session(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)
	for _, msg := range msgs {
		nameSlices := strings.Split(msg.SourceName, ";")
		if len(nameSlices) == 3 {
			stmt := fmt.Sprintf(deleteAccountCypher, msg.TableName, nameSlices[1])
			if _, err := session.Run(ctx, stmt, map[string]any{"cq_source_name": nameSlices[0], "cq_sync_time": msg.SyncTime.Truncate(time.Microsecond), "account_id": nameSlices[2]}); err != nil {
				return err
			}
			// fmt.Println(stmt)
		} else {
			stmt := fmt.Sprintf(deleteCypher, msg.TableName)
			if _, err := session.Run(ctx, stmt, map[string]any{"cq_source_name": nameSlices[0], "cq_sync_time": msg.SyncTime.Truncate(time.Microsecond)}); err != nil {
				return err
			}
			// fmt.Println(stmt)
		}

	}
	return session.Close(ctx)
}
