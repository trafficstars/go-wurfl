package wurfl

var (
	agentsList = []string{
		"Mozilla/5.0 (Windows NT 5.1; rv:15.0) Gecko/20120911 Firefox/15.1 PaleMoon/15.1",
		"Mozilla/5.0 (Amiga; U; AmigaOS 1.3; en; rv:1.8.1.19) Gecko/20081204 SeaMonkey/1.1.14",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_8_1) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1200.0 Iron/21.0.1200.0 Safari/537.1",
		"Mozilla/5.0 (Linux; Android 5.0; SAMSUNG SM-G925 Build/LRX21V) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/4.0 Chrome/44.0.2403.133 Mobile Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/534.24 (KHTML, like Gecko) RockMelt/0.9.58.390 Chrome/11.0.696.71 Safari/534.24",
		"Mozilla/5.0 (Windows; U; Win 9x 4.90; en-US; rv:1.8.1.8pre) Gecko/20071015 Firefox/2.0.0.7 Navigator/9.0",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; ko; rv:1.9.2.16) Gecko/20110319 Firefox/3.6.16 ( .NET CLR 3.5.30729; .NET4.0E)",
		"Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.107 Safari/535.1",
		"Opera/5.0 (Ubuntu; U; Windows NT 6.1; es; rv:1.9.2.13) Gecko/20101203 Firefox/3.6.13",
		"Opera/9.80 (X11; Linux x86_64; U; Ubuntu/10.10 (maverick); pl) Presto/2.7.62 Version/11.01",
	}

	keywords = []string{
		"aac",
		"accept_third_party_cookie",
		"access_key_support",
		"ajax_manipulate_css",
		"ajax_manipulate_dom",
		"ajax_preferred_geoloc_api",
		"ajax_support_events",
		"ajax_support_event_listener",
		"ajax_support_getelementbyid",
		"ajax_support_inner_html",
		"ajax_support_javascript",
		"ajax_xhr_type",
		"amr",
		"ascii_support",
		"au",
		"awb",
		"basic_authentication_support",
		"bmp",
		"brand_name",
		"break_list_of_links_with_br_element_recommended",
		"built_in_back_button_support",
		"built_in_camera",
		"built_in_recorder",
		"callericon",
		"canvas_support",
		"can_assign_phone_number",
		"can_skip_aligned_link_row",
		"card_title_support",
		"chtml_can_display_images_and_text_on_same_line",
		"chtml_displays_image_in_center",
		"chtml_display_accesskey",
		"chtml_make_phone_call_string",
		"chtml_table_support",
		"colors",
		"columns",
		"compactmidi",
		"connectionless_cache_operation",
		"connectionless_service_indication",
		"connectionless_service_load",
		"connectionoriented_confirmed_cache_operation",
		"connectionoriented_confirmed_service_indication",
		"connectionoriented_confirmed_service_load",
		"connectionoriented_unconfirmed_cache_operation",
		"connectionoriented_unconfirmed_service_indication",
		"connectionoriented_unconfirmed_service_load",
		"cookie_support",
		"css_border_image",
		"css_gradient",
		"css_gradient_linear",
		"css_rounded_corners",
		"css_spriting",
		"css_supports_width_as_percentage",
		"deck_prefetch_support",
		"density_class",
		"device_claims_web_support",
		"device_os",
		"device_os_version",
		"digiplug",
		"directdownload_support",
		"doja_1_0",
		"doja_1_5",
		"doja_2_0",
		"doja_2_1",
		"doja_2_2",
		"doja_3_0",
		"doja_3_5",
		"doja_4_0",
		"downloadfun_support",
		"dual_orientation",
		"elective_forms_recommended",
		"emoji",
		"emptyok",
		"empty_option_value_support",
		"ems",
		"ems_imelody",
		"ems_odi",
		"ems_upi",
		"ems_variablesizedpictures",
		"ems_version",
		"epoc_bmp",
		"evrc",
		"expiration_date",
		"flash_lite_version",
		"fl_browser",
		"fl_screensaver",
		"fl_standalone",
		"fl_sub_lcd",
		"fl_wallpaper",
		"full_flash_support",
		"gif",
		"gif_animated",
		"gprtf",
		"greyscale",
		"handheldfriendly",
		"has_cellular_radio",
		"has_qwerty_keyboard",
		"hinted_progressive_download",
		"html_preferred_dtd",
		"html_web_3_2",
		"html_web_4_0",
		"html_wi_imode_compact_generic",
		"html_wi_imode_htmlx_1",
		"html_wi_imode_htmlx_1_1",
		"html_wi_imode_html_1",
		"html_wi_imode_html_2",
		"html_wi_imode_html_3",
		"html_wi_imode_html_4",
		"html_wi_imode_html_5",
		"html_wi_oma_xhtmlmp_1_0",
		"html_wi_w3_xhtmlbasic",
		"https_support",
		"icons_on_menu_items_support",
		"image_as_link_support",
		"image_inlining",
		"imelody",
		"imode_region",
		"inline_support",
		"insert_br_element_after_widget_recommended",
		"iso8859_support",
		"is_bot",
		"is_console",
		"is_google_glass",
		"is_sencha_touch_ok",
		"is_smarttv",
		"is_tablet",
		"is_transcoder",
		"is_wireless_device",
		"j2me_3dapi",
		"j2me_3gpp",
		"j2me_aac",
		"j2me_amr",
		"j2me_au",
		"j2me_audio_capture_enabled",
		"j2me_bits_per_pixel",
		"j2me_bmp",
		"j2me_bmp3",
		"j2me_btapi",
		"j2me_canvas_height",
		"j2me_canvas_width",
		"j2me_capture_image_formats",
		"j2me_cldc_1_0",
		"j2me_cldc_1_1",
		"j2me_clear_key_code",
		"j2me_datefield_broken",
		"j2me_datefield_no_accepts_null_date",
		"j2me_gif",
		"j2me_gif89a",
		"j2me_h263",
		"j2me_heap_size",
		"j2me_http",
		"j2me_https",
		"j2me_imelody",
		"j2me_jpg",
		"j2me_jtwi",
		"j2me_left_softkey_code",
		"j2me_locapi",
		"j2me_max_jar_size",
		"j2me_max_record_store_size",
		"j2me_middle_softkey_code",
		"j2me_midi",
		"j2me_midp_1_0",
		"j2me_midp_2_0",
		"j2me_mmapi_1_0",
		"j2me_mmapi_1_1",
		"j2me_motorola_lwt",
		"j2me_mp3",
		"j2me_mp4",
		"j2me_mpeg4",
		"j2me_nokia_ui",
		"j2me_photo_capture_enabled",
		"j2me_png",
		"j2me_real8",
		"j2me_realaudio",
		"j2me_realmedia",
		"j2me_realvideo",
		"j2me_return_key_code",
		"j2me_right_softkey_code",
		"j2me_rmf",
		"j2me_screen_height",
		"j2me_screen_width",
		"j2me_select_key_code",
		"j2me_serial",
		"j2me_siemens_color_game",
		"j2me_siemens_extension",
		"j2me_socket",
		"j2me_storage_size",
		"j2me_svgt",
		"j2me_udp",
		"j2me_video_capture_enabled",
		"j2me_wav",
		"j2me_wbmp",
		"j2me_wma",
		"j2me_wmapi_1_0",
		"j2me_wmapi_1_1",
		"j2me_wmapi_2_0",
		"j2me_xmf",
		"jpeg_2000",
		"jpeg_xr",
		"jpg",
		"jqm_grade",
		"largeoperatorlogo",
		"manufacturer_name",
		"marketing_name",
		"max_data_rate",
		"max_deck_size",
		"max_image_height",
		"max_image_width",
		"max_length_of_password",
		"max_length_of_username",
		"max_no_of_bookmarks",
		"max_no_of_connection_settings",
		"max_object_size",
		"max_url_length_bookmark",
		"max_url_length_cached_page",
		"max_url_length_homepage",
		"max_url_length_in_requests",
		"menu_with_list_of_links_recommended",
		"menu_with_select_element_recommended",
		"midi_monophonic",
		"midi_polyphonic",
		"mld",
		"mmf",
		"mms_3gpp",
		"mms_3gpp2",
		"mms_amr",
		"mms_bmp",
		"mms_evrc",
		"mms_gif_animated",
		"mms_gif_static",
		"mms_jad",
		"mms_jar",
		"mms_jpeg_baseline",
		"mms_jpeg_progressive",
		"mms_max_frame_rate",
		"mms_max_height",
		"mms_max_size",
		"mms_max_width",
		"mms_midi_monophonic",
		"mms_midi_polyphonic",
		"mms_midi_polyphonic_voices",
		"mms_mmf",
		"mms_mp3",
		"mms_mp4",
		"mms_nokia_3dscreensaver",
		"mms_nokia_operatorlogo",
		"mms_nokia_ringingtone",
		"mms_nokia_wallpaper",
		"mms_ota_bitmap",
		"mms_png",
		"mms_qcelp",
		"mms_rmf",
		"mms_spmidi",
		"mms_symbian_install",
		"mms_vcalendar",
		"mms_vcard",
		"mms_video",
		"mms_wav",
		"mms_wbmp",
		"mms_wbxml",
		"mms_wml",
		"mms_wmlc",
		"mms_xmf",
		"mobileoptimized",
		"mobile_browser",
		"mobile_browser_version",
		"model_extra_info",
		"model_name",
		"mp3",
		"multipart_support",
		"nfc_support",
		"nokiaring",
		"nokiavcal",
		"nokiavcard",
		"nokia_edition",
		"nokia_feature_pack",
		"nokia_ringtone",
		"nokia_series",
		"nokia_voice_call",
		"numbered_menus",
		"oma_support",
		"oma_v_1_0_combined_delivery",
		"oma_v_1_0_forwardlock",
		"oma_v_1_0_separate_delivery",
		"operatorlogo",
		"opwv_wml_extensions_support",
		"opwv_xhtml_extensions_support",
		"panasonic",
		"pdf_support",
		"phone_id_provided",
		"physical_screen_height",
		"physical_screen_width",
		"picture",
		"picturemessage",
		"picture_bmp",
		"picture_colors",
		"picture_df_size_limit",
		"picture_directdownload_size_limit",
		"picture_gif",
		"picture_greyscale",
		"picture_inline_size_limit",
		"picture_jpg",
		"picture_max_height",
		"picture_max_width",
		"picture_oma_size_limit",
		"picture_png",
		"picture_preferred_height",
		"picture_preferred_width",
		"picture_resize",
		"picture_wbmp",
		"playback_3g2",
		"playback_3gpp",
		"playback_acodec_aac",
		"playback_acodec_amr",
		"playback_acodec_qcelp",
		"playback_df_size_limit",
		"playback_directdownload_size_limit",
		"playback_inline_size_limit",
		"playback_mov",
		"playback_mp4",
		"playback_oma_size_limit",
		"playback_real_media",
		"playback_vcodec_h263_0",
		"playback_vcodec_h263_3",
		"playback_vcodec_h264_bp",
		"playback_vcodec_mpeg4_asp",
		"playback_vcodec_mpeg4_sp",
		"playback_wmv",
		"png",
		"pointing_method",
		"post_method_support",
		"preferred_markup",
		"progressive_download",
		"proportional_font",
		"qcelp",
		"receiver",
		"release_date",
		"resolution_height",
		"resolution_width",
		"ringtone",
		"ringtone_3gpp",
		"ringtone_aac",
		"ringtone_amr",
		"ringtone_awb",
		"ringtone_compactmidi",
		"ringtone_df_size_limit",
		"ringtone_digiplug",
		"ringtone_directdownload_size_limit",
		"ringtone_imelody",
		"ringtone_inline_size_limit",
		"ringtone_midi_monophonic",
		"ringtone_midi_polyphonic",
		"ringtone_mmf",
		"ringtone_mp3",
		"ringtone_oma_size_limit",
		"ringtone_qcelp",
		"ringtone_rmf",
		"ringtone_spmidi",
		"ringtone_voices",
		"ringtone_wav",
		"ringtone_xmf",
		"rmf",
		"rows",
		"rss_support",
		"sagem_v1",
		"sagem_v2",
		"sckl_groupgraphic",
		"sckl_operatorlogo",
		"sckl_ringtone",
		"sckl_vcalendar",
		"sckl_vcard",
		"screensaver",
		"screensaver_bmp",
		"screensaver_colors",
		"screensaver_df_size_limit",
		"screensaver_directdownload_size_limit",
		"screensaver_gif",
		"screensaver_greyscale",
		"screensaver_inline_size_limit",
		"screensaver_jpg",
		"screensaver_max_height",
		"screensaver_max_width",
		"screensaver_oma_size_limit",
		"screensaver_png",
		"screensaver_preferred_height",
		"screensaver_preferred_width",
		"screensaver_resize",
		"screensaver_wbmp",
		"sdio",
		"sender",
		"siemens_logo_height",
		"siemens_logo_width",
		"siemens_ota",
		"siemens_screensaver_height",
		"siemens_screensaver_width",
		"smf",
		"sms_enabled",
		"softkey_support",
		"sp_midi",
		"streaming_3g2",
		"streaming_3gpp",
		"streaming_acodec_aac",
		"streaming_acodec_amr",
		"streaming_flv",
		"streaming_mov",
		"streaming_mp4",
		"streaming_preferred_http_protocol",
		"streaming_preferred_protocol",
		"streaming_real_media",
		"streaming_vcodec_h263_0",
		"streaming_vcodec_h263_3",
		"streaming_vcodec_h264_bp",
		"streaming_vcodec_mpeg4_asp",
		"streaming_vcodec_mpeg4_sp",
		"streaming_video",
		"streaming_video_size_limit",
		"streaming_wmv",
		"svgt_1_1",
		"svgt_1_1_plus",
		"table_support",
		"text_imelody",
		"tiff",
		"times_square_mode_support",
		"time_to_live_support",
		"total_cache_disable_support",
		"transcoder_ua_header",
		"transparent_png_alpha",
		"transparent_png_index",
		"uaprof",
		"uaprof2",
		"uaprof3",
		"unique",
		"ununiqueness_handler",
		"utf8_support",
		"ux_full_desktop",
		"video",
		"viewport_initial_scale",
		"viewport_maximum_scale",
		"viewport_minimum_scale",
		"viewport_supported",
		"viewport_userscalable",
		"viewport_width",
		"voices",
		"voicexml",
		"vpn",
		"wallpaper",
		"wallpaper_bmp",
		"wallpaper_colors",
		"wallpaper_df_size_limit",
		"wallpaper_directdownload_size_limit",
		"wallpaper_gif",
		"wallpaper_greyscale",
		"wallpaper_inline_size_limit",
		"wallpaper_jpg",
		"wallpaper_max_height",
		"wallpaper_max_width",
		"wallpaper_oma_size_limit",
		"wallpaper_png",
		"wallpaper_preferred_height",
		"wallpaper_preferred_width",
		"wallpaper_resize",
		"wallpaper_tiff",
		"wallpaper_wbmp",
		"wap_push_support",
		"wav",
		"wbmp",
		"webp_lossless_support",
		"webp_lossy_support",
		"wifi",
		"wizards_recommended",
		"wml_1_1",
		"wml_1_2",
		"wml_1_3",
		"wml_can_display_images_and_text_on_same_line",
		"wml_displays_image_in_center",
		"wml_make_phone_call_string",
		"wrap_mode_support",
		"wta_misc",
		"wta_pdc",
		"wta_phonebook",
		"wta_voice_call",
		"xhtmlmp_preferred_mime_type",
		"xhtml_allows_disabled_form_elements",
		"xhtml_autoexpand_select",
		"xhtml_avoid_accesskeys",
		"xhtml_can_embed_video",
		"xhtml_display_accesskey",
		"xhtml_document_title_support",
		"xhtml_file_upload",
		"xhtml_format_as_attribute",
		"xhtml_format_as_css_property",
		"xhtml_honors_bgcolor",
		"xhtml_make_phone_call_string",
		"xhtml_marquee_as_css_property",
		"xhtml_nowrap_mode",
		"xhtml_preferred_charset",
		"xhtml_readable_background_color1",
		"xhtml_readable_background_color2",
		"xhtml_select_as_dropdown",
		"xhtml_select_as_popup",
		"xhtml_select_as_radiobutton",
		"xhtml_send_mms_string",
		"xhtml_send_sms_string",
		"xhtml_supports_css_cell_table_coloring",
		"xhtml_supports_forms_in_table",
		"xhtml_supports_iframe",
		"xhtml_supports_inline_input",
		"xhtml_supports_invisible_text",
		"xhtml_supports_monospace_font",
		"xhtml_supports_table_for_layout",
		"xhtml_support_level",
		"xhtml_support_wml2_namespace",
		"xhtml_table_support",
		"xmf",
	}
)