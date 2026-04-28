components {
  id: "bin"
  component: "/assets/bin/bin.script"
}
components {
  id: "muzzle"
  component: "/assets/bin/muzzle_wheel.particlefx"
  position {
    x: -23.0
    y: -23.0
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"bin\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/props.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"poop\"\n"
  "mask: \"gears\"\n"
  "mask: \"building\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 22.5\n"
  "  data: 22.5\n"
  "  data: 32.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "rail_sound"
  type: "sound"
  data: "sound: \"/assets/bin/train_wheels.ogg\"\n"
  "looping: 1\n"
  ""
}
embedded_components {
  id: "fart"
  type: "sound"
  data: "sound: \"/assets/fart.ogg\"\n"
  ""
}
embedded_components {
  id: "metal"
  type: "sound"
  data: "sound: \"/assets/metal.ogg\"\n"
  ""
}
