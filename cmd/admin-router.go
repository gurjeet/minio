/*
 * Minio Cloud Storage, (C) 2016 Minio, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import router "github.com/gorilla/mux"

// adminAPIHandlers provides HTTP handlers for Minio admin API.
type adminAPIHandlers struct {
}

// registerAdminRouter - Add handler functions for each service REST API routes.
func registerAdminRouter(mux *router.Router) {

	adminAPI := adminAPIHandlers{}
	// Admin router
	adminRouter := mux.NewRoute().PathPrefix("/").Subrouter()

	/// Admin operations

	// Service status
	adminRouter.Methods("GET").Queries("service", "").Headers(minioAdminOpHeader, "status").HandlerFunc(adminAPI.ServiceStatusHandler)
	// Service stop
	adminRouter.Methods("POST").Queries("service", "").Headers(minioAdminOpHeader, "stop").HandlerFunc(adminAPI.ServiceStopHandler)
	// Service restart
	adminRouter.Methods("POST").Queries("service", "").Headers(minioAdminOpHeader, "restart").HandlerFunc(adminAPI.ServiceRestartHandler)
}
