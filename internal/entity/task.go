package entity
// Task represents the task for this application
//
// swagger:model Task
type Task struct {
    // The ID for this Task
    //
    // required: true
    Id uint64 `json:"id"`
    // The title for this Task
    //
    // required: true
    Title string `json:"title"`
    // The Completed represents if a task is completed or not
    //
    // required: true
    Completed bool `json:"completed"`
}