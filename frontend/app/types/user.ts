export interface User {
  id: number;
  username: string;
  email: string;
  display_name: { String: string; Valid: boolean };
  bio: { String: string; Valid: boolean };
  avatar_url: { String: string; Valid: boolean };
  created_at: string;
}
