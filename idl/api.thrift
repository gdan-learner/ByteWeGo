namespace go api

struct User {
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}

struct CheckUserRequest {
    1: string username (vt.max_size = "32")
    2: string password (vt.max_size = "32")
}

struct CheckUserResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct CreateUserRequest {
    1: string username (vt.max_size = "32")
    2: string password (vt.max_size = "32")
}

struct CreateUserResponse {
    1: i32 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct GetUserRequest {
    1: i64 user_id
    2: string token
}

struct GetUserResponse {
    1: i32 status_code
    2: string status_msg
    3: User user
}

service ApiService {
    CheckUserResponse CheckUser(1: CheckUserRequest req) (api.post="/douyin/user/login")
    CreateUserResponse CreateUser(1: CreateUserRequest req) (api.post="/douyin/user/register")
    GetUserResponse GetUser(1: GetUserRequest req) (api.get="/douyin/user")
}