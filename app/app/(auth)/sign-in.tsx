import { StyleSheet } from "react-native";
import { useHello } from "../../api/hooks/useHello";

import EditScreenInfo from "../../components/EditScreenInfo";
import { Text, View } from "../../components/Themed";
import { useAuth } from "../../context/auth";

export default function TabOneScreen() {
  const { isLoading, data } = useHello();

  console.log(isLoading, data);
  const { signIn } = useAuth();

  return (
    <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
      <Text onPress={() => signIn()}>Sign In</Text>
    </View>
  );
}
