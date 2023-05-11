package open_im_sdk_callback

type OnListenerForService interface {
	OnGroupApplicationAdded(groupApplication string)

	OnGroupApplicationAccepted(groupApplication string)

	OnFriendApplicationAdded(friendApplication string)

	OnFriendApplicationAccepted(groupApplication string)

	OnRecvNewMessage(message string)
}
