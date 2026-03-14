package flow

import "context"

type TeamService struct{ http *httpClient }

func (s *TeamService) GetTeam(ctx context.Context) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.http.get(ctx, "/merchants/me/team", nil, &result)
	return result, err
}

func (s *TeamService) GetMyRole(ctx context.Context) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.http.get(ctx, "/merchants/me/team/my-role", nil, &result)
	return result, err
}

func (s *TeamService) Invite(ctx context.Context, email, role string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.http.post(ctx, "/merchants/me/team/invite", map[string]string{"email": email, "role": role}, &result)
	return result, err
}

func (s *TeamService) RevokeInvite(ctx context.Context, inviteID string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.http.post(ctx, "/merchants/me/team/invite/"+inviteID+"/revoke", nil, &result)
	return result, err
}

func (s *TeamService) UpdateRole(ctx context.Context, memberID, role string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := s.http.patch(ctx, "/merchants/me/team/"+memberID+"/role", map[string]string{"role": role}, &result)
	return result, err
}

func (s *TeamService) Remove(ctx context.Context, memberID string) error {
	return s.http.delete(ctx, "/merchants/me/team/"+memberID)
}
