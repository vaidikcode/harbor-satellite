# Harbor Satellite Eventing System: Supported Events

## Event Structure

| Field     | Type                   | Description                                 |
|-----------|------------------------|---------------------------------------------|
| Type      | string                 | Name of the event (e.g., CONFIG_UPDATED)    |
| Timestamp | time.Time              | When the event occurred                     |
| Source    | string                 | Component that generated the event          |
| Payload   | map[string]interface{} | Event-specific data (optional, flexible)    |

---

## Supported Events

### CONFIG_UPDATED
- **Description:** When new settings/config are fetched from Ground Control.
- **Example Payload:**
  ```json
  {
    "configVersion": "v1.2.3"
  }
  ```

### REGISTRY_STARTED
- **Description:** Emitted when the registry is successfully started by the system.
- **Payload:**
  - `address` (string): The address the registry is listening on
  - `port` (string): The port the registry is listening on
- **Example Payload:**
  ```json
  {
    "address": "127.0.0.1",
    "port": "8585"
  }
  ```

---

> More events can be added following this structure. See `internal/eventbus/event.go` for the event definition. 