export interface Post {
  id: number;
  author_id: number;
  content: {
    String: string;
    Valid: boolean;
  };
  reply_count: number;
  like_count: number;
  repost_count: number;
  created_at: string;
  updated_at: string;
}

export interface ApiResponse<T> {
  data?: T;
  error?: string;
}
